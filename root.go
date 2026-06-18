package cmd

import (
	"os"

	"github.com/velonetics/lura/v2/config"
	"github.com/velonetics/lura/v2/core"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var IsTTY = isatty.IsTerminal(os.Stderr.Fd())

var (
	cfgFile              string
	debug                int
	port                 int
	usageDisable         bool
	checkGinRoutes       bool
	checkDebug           int
	lintCurrentSchema    bool
	lintCustomSchemaPath string
	lintNoNetwork        bool
	rawEmbedSchema       string
	rulesToExclude       string
	rulesToExcludePath   string
	severitiesToInclude  = "CRITICAL,HIGH,MEDIUM,LOW"
	formatTmpl           string
	parser               config.Parser
	run                  func(config.ServiceConfig)

	goSum           = "./go.sum"
	goVersion       = core.GoVersion
	libcVersion     = core.GlibcVersion
	checkDumpPrefix = "\t"
	gogetEnabled    = false

	DefaultRoot    Root
	RootCommand    Command
	RunCommand     Command
	CheckCommand   Command
	PluginCommand  Command
	VersionCommand Command
	AuditCommand   Command

	rootCmd = &cobra.Command{
		Use:   "velonetics",
		Short: "Velonetics is a high-performance API gateway that helps you publish, secure, control, and monitor your services",
	}

	checkCmd = &cobra.Command{
		Use:     "check",
		Short:   "Validates that the configuration file is valid.",
		Long:    "Validates that the active configuration file has a valid syntax to run the service.\nChange the configuration file by using the --config flag",
		Run:     checkFunc,
		Aliases: []string{"validate"},
		Example: "velonetics check -d -l -c config.json",
	}

	runCmd = &cobra.Command{
		Use:     "run",
		Short:   "Runs the Velonetics server.",
		Long:    "Runs the Velonetics server.",
		Run:     runFunc,
		Example: "velonetics run -d -c config.json",
	}

	pluginCmd = &cobra.Command{
		Use:     "check-plugin",
		Short:   "Checks your plugin dependencies are compatible.",
		Long:    "Checks your plugin dependencies are compatible and proposes commands to update your dependencies.",
		Run:     pluginFunc,
		Example: "velonetics check-plugin -g 1.19.0 -s ./go.sum -f",
	}

	versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Shows Velonetics version.",
		Long:    "Shows Velonetics version.",
		Run:     versionFunc,
		Example: "velonetics version",
	}

	auditCmd = &cobra.Command{
		Use:     "audit",
		Short:   "Audits a Velonetics configuration.",
		Long:    "Audits a Velonetics configuration.",
		Run:     auditFunc,
		Example: "velonetics audit -i 1.1.1,1.1.2 -s CRITICAL -c velonetics.json",
	}
)

func init() {
	cfgFlag := StringFlagBuilder(&cfgFile, "config", "c", "", "Path to the configuration file")
	debugFlag := CountFlagBuilder(&debug, "debug", "d", "Enables the debug endpoint")
	RootCommand = NewCommand(rootCmd)
	RootCommand.Cmd.SetHelpTemplate(logoBanner + "Version: " + core.VeloneticsVersion + "\n\n" + rootCmd.HelpTemplate())

	ginRoutesFlag := BoolFlagBuilder(&checkGinRoutes, "test-gin-routes", "t", false, "Tests the endpoint patterns against a real gin router on the selected port")
	prefixFlag := StringFlagBuilder(&checkDumpPrefix, "indent", "i", checkDumpPrefix, "Indentation of the check dump")
	lintCurrentSchemaFlag := BoolFlagBuilder(&lintCurrentSchema, "lint", "l", lintCurrentSchema, "Enables the linting against the official Velonetics online JSON schema")
	lintCustomSchemaFlag := StringFlagBuilder(&lintCustomSchemaPath, "lint-schema", "s", lintCustomSchemaPath, "Lint against a custom schema path or URL")
	lintNoNetworkFlag := BoolFlagBuilder(&lintNoNetwork, "lint-no-network", "n", lintNoNetwork, "Lint against the builtin Velonetics JSON schema, no network is required")
	checkDebugFlag := CountFlagBuilder(&checkDebug, "debug", "d", "Information about how Velonetics is interpreting your configuration file")
	CheckCommand = NewCommand(checkCmd, cfgFlag, checkDebugFlag, ginRoutesFlag, prefixFlag, lintCurrentSchemaFlag, lintCustomSchemaFlag, lintNoNetworkFlag)
	CheckCommand.AddConstraint(MutuallyExclusive("lint", "lint-no-network", "lint-schema"))

	portFlag := IntFlagBuilder(&port, "port", "p", 0, "Listening port for the http service")
	usageDisableFlag := BoolFlagBuilder(&usageDisable, "usage-disable", "", false, "Disables anonymous usage reporting")
	RunCommand = NewCommand(runCmd, cfgFlag, debugFlag, portFlag, usageDisableFlag)

	goSumFlag := StringFlagBuilder(&goSum, "sum", "s", goSum, "Path to the go.sum file to analyze")
	goVersionFlag := StringFlagBuilder(&goVersion, "go", "g", goVersion, "The version of the go compiler used for your plugin")
	libcVersionFlag := StringFlagBuilder(&libcVersion, "libc", "l", "", "Version of the libc library used")
	gogetFlag := BoolFlagBuilder(&gogetEnabled, "format", "f", false, "Shows fix commands to update your dependencies")
	PluginCommand = NewCommand(pluginCmd, goSumFlag, goVersionFlag, libcVersionFlag, gogetFlag)

	rulesToExcludeFlag := StringFlagBuilder(&rulesToExclude, "ignore", "i", rulesToExclude, "List of rules to ignore (comma-separated, no spaces)")
	severitiesToIncludeFlag := StringFlagBuilder(&severitiesToInclude, "severity", "s", severitiesToInclude, "List of severities to include (comma-separated, no spaces)")
	pathToRulesToExcludeFlag := StringFlagBuilder(&rulesToExcludePath, "ignore-file", "I", rulesToExcludePath, "Path to a text-plain file containing the list of rules to exclude")
	formatFlag := StringFlagBuilder(&formatTmpl, "format", "f", formatTmpl, "Inline go template to render the results")
	AuditCommand = NewCommand(auditCmd, cfgFlag, rulesToExcludeFlag, severitiesToIncludeFlag, pathToRulesToExcludeFlag, formatFlag)

	VersionCommand = NewCommand(versionCmd)

	DefaultRoot = NewRoot(RootCommand, CheckCommand, RunCommand, PluginCommand, VersionCommand, AuditCommand)
}

const logoBanner = `  velonetics
  API Gateway

`
