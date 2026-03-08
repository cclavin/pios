/**
 * PIOS Excuse Generator Widget
 * Lightweight, zero-dependency, CWV-friendly embeddable widget.
 */
(function () {
    'use strict';

    // Look for the embedding container
    const containerId = 'pios-excuse-widget';
    const container = document.getElementById(containerId);

    if (!container) {
        console.warn(`[PIOS Widget] Container #${containerId} not found.`);
        return;
    }

    // Default configuration (can be overridden via data attributes)
    const config = {
        buttonText: container.getAttribute('data-button-text') || 'Generate Excuse',
        theme: container.getAttribute('data-theme') || 'light', // 'light' or 'dark'
        heading: container.getAttribute('data-heading') || 'Random Excuse Generator'
    };

    // The dataset
    const excuses = [
        "A background task consumed all my memory.",
        "The compiler deleted my source code by mistake.",
        "My cat walked across the keyboard and deployed to production.",
        "I thought that feature was out-of-scope for v1.",
        "The third-party API is experiencing intermittent timeouts.",
        "A stray cosmic ray flipped a bit in the server RAM.",
        "I was testing a hypothesis in production.",
        "The PM told me it wasn't a priority.",
        "My local environment is perfectly fine, it must be the server.",
        "I'm waiting on another team to finish their microservice.",
        "I had to clear my cache, and now everything is broken."
    ];

    // Inject styles
    const styleId = 'pios-excuse-styles';
    if (!document.getElementById(styleId)) {
        const style = document.createElement('style');
        style.id = styleId;
        // Prefix all classes with pios-excuse- to prevent CSS bleeding
        style.textContent = `
            #${containerId} {
                font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
                box-sizing: border-box;
                display: flex;
                flex-direction: column;
                align-items: center;
                text-align: center;
                padding: 1.5rem;
                border-radius: 8px;
                max-width: 400px;
                margin: 0 auto;
                box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
                transition: background-color 0.3s ease, color 0.3s ease;
                /* CWV: Pre-allocate min-height to prevent layout shifts when text changes */
                min-height: 200px; 
            }
            #${containerId}[data-theme="light"] {
                background-color: #ffffff;
                color: #333333;
                border: 1px solid #e5e7eb;
            }
            #${containerId}[data-theme="dark"] {
                background-color: #1f2937;
                color: #f9fafb;
                border: 1px solid #374151;
            }
            .pios-excuse-heading {
                margin: 0 0 1rem 0;
                font-size: 1.25rem;
                font-weight: 600;
            }
            .pios-excuse-display {
                flex-grow: 1;
                display: flex;
                align-items: center;
                justify-content: center;
                font-size: 1.1rem;
                font-style: italic;
                margin-bottom: 1.5rem;
                min-height: 3rem; /* Prevents button jumping */
            }
            .pios-excuse-btn {
                appearance: none;
                background-color: #3b82f6;
                color: white;
                border: none;
                padding: 0.5rem 1rem;
                font-size: 1rem;
                font-weight: 500;
                border-radius: 6px;
                cursor: pointer;
                transition: background-color 0.2s ease;
            }
            .pios-excuse-btn:hover {
                background-color: #2563eb;
            }
            .pios-excuse-btn:active {
                background-color: #1d4ed8;
            }
            #${containerId}[data-theme="dark"] .pios-excuse-btn {
                background-color: #4f46e5;
            }
            #${containerId}[data-theme="dark"] .pios-excuse-btn:hover {
                background-color: #4338ca;
            }
        `;
        document.head.appendChild(style);
    }

    // Ensure theme state is on container
    container.setAttribute('data-theme', config.theme);

    // Build DOM
    const heading = document.createElement('h3');
    heading.className = 'pios-excuse-heading';
    heading.textContent = config.heading;

    const display = document.createElement('div');
    display.className = 'pios-excuse-display';
    display.textContent = "Click the button for an excuse!";

    const button = document.createElement('button');
    button.className = 'pios-excuse-btn';
    button.textContent = config.buttonText;

    // Attach event listener
    button.addEventListener('click', () => {
        const randomIndex = Math.floor(Math.random() * excuses.length);
        const selected = excuses[randomIndex];
        display.textContent = `"${selected}"`;

        // Lightweight animation
        display.style.opacity = '0';
        setTimeout(() => {
            display.style.transition = 'opacity 0.2s ease-in';
            display.style.opacity = '1';
        }, 50);
    });

    // Inject into container
    container.appendChild(heading);
    container.appendChild(display);
    container.appendChild(button);

})();
