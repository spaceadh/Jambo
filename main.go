package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spaceadh/Jambo/repl"
	"github.com/spaceadh/Jambo/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title = styles.TitleStyle.
		Render(`


		╋╋┏┳━━━┳━┓┏━┳━━┓┏━━━┓
		╋╋┃┃┏━┓┃┃┗┛┃┃┏┓┃┃┏━┓┃
		╋╋┃┃┃╋┃┃┏┓┏┓┃┗┛┗┫┃╋┃┃
		┏┓┃┃┗━┛┃┃┃┃┃┃┏━┓┃┃╋┃┃
		┃┗┛┃┏━┓┃┃┃┃┃┃┗━┛┃┗━┛┃
		┗━━┻┛╋┗┻┛┗┛┗┻━━━┻━━━┛

`)
	Version = styles.VersionStyle.Render("v0.5.1")
	Author  = styles.AuthorStyle.Render("by Alvin Wachira")
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
	Help    = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia Jambo:
	%s: Kuanza programu ya Jambo
	%s: Kuendesha faili la Jambo
	%s: Kusoma nyaraka za Jambo
	%s: Kufahamu toleo la Jambo
`,
		styles.HelpStyle.Bold(true).Render("jambo"),
		styles.HelpStyle.Bold(true).Render("jambo jinaLaFile.jb"),
		styles.HelpStyle.Bold(true).Render("jambo --nyaraka"),
		styles.HelpStyle.Bold(true).Render("jambo --toleo")))
)

func main() {

	args := os.Args
	if len(args) < 2 {

		help := styles.HelpStyle.Render("💡 Tumia exit() au toka() kuondoka")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		os.Exit(0)
	}

	if len(args) == 2 {

		switch args[1] {
		case "msaada", "-msaada", "--msaada", "help", "-help", "--help", "-h":
			fmt.Println(Help)
			os.Exit(0)
		case "version", "-version", "--version", "-v", "v", "--toleo", "-toleo":
			fmt.Println(NewLogo)
			os.Exit(0)
		case "-docs", "--docs", "-nyaraka", "--nyaraka":
			repl.Docs()
			os.Exit(0)
		}

		file := args[1]

		if strings.HasSuffix(file, "jb") || strings.HasSuffix(file, ".sw") {
			contents, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: Jambo halipati faili liitwalo : ", args[1]))
				os.Exit(0)
			}

			repl.Read(string(contents))
		} else {
			fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "sii faili sahihi. Tumia faili la '.jb' au '.sw'"))
			os.Exit(0)
		}

	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: Operesheni imeshindikana boss."))
		fmt.Println(Help)
		os.Exit(0)
	}
}
