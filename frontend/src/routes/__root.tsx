import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'
import { Auth } from '../components/Auth'  // ✅ Import Auth

const RootLayout = () => (
    <>
        <nav className="bg-white shadow">
            <div className="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
                <div className="flex gap-4 items-center">
                    <h1 className="text-xl font-bold">Deploy Watch</h1>
                    <Link to="/" className="[&.active]:font-bold text-sm">
                        Home
                    </Link>
                    <Link to="/about" className="[&.active]:font-bold text-sm">
                        About
                    </Link>
                </div>
                <Auth />  {/* ✅ Auth component here */}
            </div>
        </nav>

        <main className="max-w-7xl mx-auto px-4 py-8">
            <Outlet />  {/* Routes render here */}
        </main>

        <TanStackRouterDevtools />
    </>
)

export const Route = createRootRoute({ component: RootLayout })