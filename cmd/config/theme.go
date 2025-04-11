package config

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func SetResponseColorTheme(text string) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0"))
	fmt.Println(style.Render(text))
	borderBottomStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFDE21")).MarginTop(1)
	borderBottom := borderBottomStyle.Render(strings.Repeat("*", 5))
	fmt.Println(borderBottom)
}