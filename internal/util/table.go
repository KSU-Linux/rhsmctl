package util

import (
    "os"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
    "github.com/charmbracelet/x/term"
)

func NewTable() *table.Table {
    var (
        blue            =   lipgloss.Color("68")
        gray            =   lipgloss.Color("249")
        lightGray       =   lipgloss.Color("253")
        headerStyle     =   lipgloss.NewStyle().Bold(true).Padding(0, 1)
        cellStyle       =   lipgloss.NewStyle().Padding(0,1)
        oddRowStyle     =   cellStyle.Foreground(gray)
        evenRowStyle    =   cellStyle.Foreground(lightGray)
        termWidth, _, _ =   term.GetSize(os.Stdout.Fd())
    )

    t := table.New().
        Border(lipgloss.RoundedBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(blue)).
        StyleFunc(func(row, col int) lipgloss.Style {
            switch {
            case row == table.HeaderRow:
                return headerStyle
            case row%2 == 0:
                return evenRowStyle
            default:
                return oddRowStyle
            }
        }).
        Width(termWidth)
    return t
}
