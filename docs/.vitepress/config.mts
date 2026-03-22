import { defineConfig } from 'vitepress'

export default defineConfig({
    base: '/pios/',
    title: "PIOS",
    description: "The AI Project Execution Contract",
    themeConfig: {
        nav: [
            { text: 'Home', link: '/' },
            { text: 'What is PIOS?', link: '/guide' },
        ],
        sidebar: [
            {
                text: 'Introduction',
                items: [
                    { text: 'What is PIOS?', link: '/guide' },
                    { text: 'Positioning', link: '/positioning' },
                    { text: 'Scope', link: '/scope' }
                ]
            },
            {
                text: 'Architecture',
                items: [
                    { text: 'Contracts', link: '/contracts' },
                    { text: 'MCP Server', link: '/mcp-server' }
                ]
            }
        ],
        socialLinks: [
            { icon: 'github', link: 'https://github.com/cclavin/pios' }
        ]
    }
})
