import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/about')({
    component: About,
})

function About() {
    return (
        <div className="max-w-4xl mx-auto space-y-8">
            <section>
                <h2 className="text-3xl font-bold mb-4">About Deploy Watch</h2>
                <p className="text-gray-700 leading-relaxed">
                    Deploy Watch is a real-time GitHub deployment monitoring application built as a learning project
                    to master Go backend development and modern web architecture patterns.
                </p>
            </section>

            <section>
                <h3 className="text-2xl font-semibold mb-3">🎯 What It Does</h3>
                <ul className="space-y-2 text-gray-700">
                    <li>• Monitors GitHub push events via webhooks</li>
                    <li>• Real-time notifications with Server-Sent Events (SSE)</li>
                    <li>• GitHub OAuth authentication</li>
                    <li>• Deploy history and analytics dashboard</li>
                    <li>• Multi-repository support</li>
                </ul>
            </section>

            <section>
                <h3 className="text-2xl font-semibold mb-3">🛠️ Technology Stack</h3>

                <div className="grid md:grid-cols-2 gap-6">
                    <div className="border rounded-lg p-4">
                        <h4 className="font-semibold text-lg mb-2">Backend</h4>
                        <ul className="space-y-1 text-gray-700">
                            <li>• Go 1.23</li>
                            <li>• Chi Router</li>
                            <li>• PostgreSQL (Neon)</li>
                            <li>• Server-Sent Events (SSE)</li>
                            <li>• OAuth 2.0</li>
                            <li>• Deployed on Fly.io</li>
                        </ul>
                    </div>

                    <div className="border rounded-lg p-4">
                        <h4 className="font-semibold text-lg mb-2">Frontend</h4>
                        <ul className="space-y-1 text-gray-700">
                            <li>• React + TypeScript</li>
                            <li>• TanStack Router</li>
                            <li>• TanStack Query</li>
                            <li>• Tailwind CSS</li>
                            <li>• Vite</li>
                            <li>• EventSource API</li>
                        </ul>
                    </div>
                </div>
            </section>

            <section>
                <h3 className="text-2xl font-semibold mb-3">🏗️ Architecture Patterns</h3>
                <div className="grid gap-4">
                    <div className="border-l-4 border-blue-500 pl-4">
                        <h4 className="font-semibold">Hub-and-Spoke Pattern (SSE Broadcaster)</h4>
                        <p className="text-gray-700 text-sm">
                            Go channels-based broadcaster for real-time user-specific event distribution
                        </p>
                    </div>

                    <div className="border-l-4 border-green-500 pl-4">
                        <h4 className="font-semibold">Repository Pattern</h4>
                        <p className="text-gray-700 text-sm">
                            Clean separation between data access and business logic
                        </p>
                    </div>

                    <div className="border-l-4 border-purple-500 pl-4">
                        <h4 className="font-semibold">Middleware Chain</h4>
                        <p className="text-gray-700 text-sm">
                            CORS, authentication, correlation IDs, and logging middleware
                        </p>
                    </div>

                    <div className="border-l-4 border-orange-500 pl-4">
                        <h4 className="font-semibold">Webhook Integration</h4>
                        <p className="text-gray-700 text-sm">
                            GitHub webhook processing with event-driven architecture
                        </p>
                    </div>
                </div>
            </section>

            <section>
                <h3 className="text-2xl font-semibold mb-3">💡 Key Learnings</h3>
                <ul className="space-y-2 text-gray-700">
                    <li>• Go concurrency patterns with channels and goroutines</li>
                    <li>• Real-time communication with Server-Sent Events</li>
                    <li>• OAuth 2.0 authentication flow implementation</li>
                    <li>• Cross-origin resource sharing (CORS) configuration</li>
                    <li>• PostgreSQL schema design and migrations</li>
                    <li>• Production deployment with Docker and Fly.io</li>
                    <li>• TypeScript type safety in React applications</li>
                </ul>
            </section>

            <section className="border-t pt-6">
                <h3 className="text-2xl font-semibold mb-3">🔗 Links</h3>

                <div className="flex gap-4">
                    <a
                        href="https://github.com/diagnosis/deploy-watch"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-blue-600 hover:underline"
                    >
                        Backend Repository →
                    </a>

                    <a
                        href="https://github.com/diagnosis"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-blue-600 hover:underline"
                    >
                        GitHub Profile →
                    </a>
                </div>
            </section>

    <section className="bg-gray-50 rounded-lg p-6">
    
    </section>
</div>
)
}