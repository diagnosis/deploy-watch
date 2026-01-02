import { useAuth } from "../hooks/useAuth";
import { appClient } from "../lib/api";

export function Auth() {
    const { user, isLoading } = useAuth();

    if (isLoading) {
        return <span className="text-gray-500">Loading...</span>;
    }

    if (!user) {
        return (
            <a
                href={appClient.getLoginUrl()}
                className="text-sm text-blue-600 hover:underline"
            >
                Login
            </a>
        );
    }

    return (
        <div className="flex items-center gap-3">
            <img
                src={user.avatar_url}
                alt={user.username}
                className="h-8 w-8 rounded-full"
            />
            <span className="text-sm text-gray-800">{user.username}</span>

            <button
                onClick={() => appClient.logout()}
                className="text-sm text-gray-500 hover:text-gray-800"
            >
                Logout
            </button>
        </div>
    );
}