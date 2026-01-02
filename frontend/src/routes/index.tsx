// routes/index.tsx
import { createFileRoute } from '@tanstack/react-router'
import {DeployList} from "../components/DeployList.tsx";

function Home() {
    return (
        <div>
            <h2 className="text-2xl font-semibold mb-4">Recent Deploys</h2>
            <DeployList />
        </div>
    )
}

export const Route = createFileRoute('/')({
    component: Home,
})