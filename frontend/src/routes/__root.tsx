import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'
import { Auth } from '../components/Auth'

const RootLayout = () => (
    <div className="min-h-screen flex flex-col bg-gray-50">
        {/* Header */}
        <nav className="bg-white shadow-sm border-b sticky top-0 z-50">
            <div className="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
                {/* Logo + Nav Links */}
                <div className="flex gap-6 items-center">
                    <Link to="/" className="flex items-center gap-2">
                        <span className="text-2xl">🚀</span>
                        <h1 className="text-xl font-bold text-gray-900">Deploy Watch</h1>
                    </Link>

                    <div className="hidden md:flex gap-4">
                        <Link
                            to="/"
                            className="text-sm text-gray-600 hover:text-gray-900 [&.active]:text-blue-600 [&.active]:font-semibold transition"
                        >
                            Home
                        </Link>
                        <Link
                            to="/stats"
                            className="text-sm text-gray-600 hover:text-gray-900 [&.active]:text-blue-600 [&.active]:font-semibold transition"
                        >
                            Stats
                        </Link>
                        <Link
                            to="/about"
                            className="text-sm text-gray-600 hover:text-gray-900 [&.active]:text-blue-600 [&.active]:font-semibold transition"
                        >
                            About
                        </Link>
                    </div>
                </div>

                {/* Auth + GitHub */}
                <div className="flex items-center gap-4">
                    <a
                        href="https://github.com/diagnosis/deploy-watch"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="hidden sm:flex items-center gap-1 text-sm text-gray-600 hover:text-gray-900 transition"
                    >
                        <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                            <path fillRule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clipRule="evenodd" />
                        </svg>
                        <span className="hidden lg:inline">GitHub</span>
                    </a>
                    <Auth />
                </div>
            </div>
        </nav>

        {/* Main Content */}
        <main className="max-w-7xl mx-auto px-4 py-8 flex-1 w-full">
            <Outlet />
        </main>

        {/* Footer */}
        <footer className="bg-white border-t mt-auto">
            <div className="max-w-7xl mx-auto px-4 py-6">
                <div className="flex flex-col md:flex-row justify-between items-center gap-4">
                    <div className="text-sm text-gray-600">
                        <p>Built with <span className="text-red-500">♥</span> using Go + React</p>
                    </div>

                    <div className="flex gap-6 text-sm">
                        <a
                            href="https://github.com/diagnosis"
                            target="_blank"
                            rel="noopener noreferrer"
                            className="text-gray-600 hover:text-gray-900 transition"
                        >
                            @diagnosis
                        </a>
                        <a
                            href="https://github.com/diagnosis/deploy-watch"
                            target="_blank"
                            rel="noopener noreferrer"
                            className="text-gray-600 hover:text-gray-900 transition"
                        >
                            Source Code
                        </a>
                        <Link
                            to="/about"
                            className="text-gray-600 hover:text-gray-900 transition"
                        >
                            Documentation
                        </Link>
                    </div>
                </div>
            </div>
        </footer>

        <TanStackRouterDevtools />
    </div>
)

export const Route = createRootRoute({ component: RootLayout })