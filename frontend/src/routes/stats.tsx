import { createFileRoute } from '@tanstack/react-router'
import {DeployStats} from "../components/DeployStats.tsx";

export const Route = createFileRoute('/stats')({
  component: Stats,
})

function Stats() {
  return (
      <div>
          <h2 className="text-2xl font-semibold mb-4">Deploy Statistics</h2>
          <DeployStats></DeployStats>
      </div>
  )
}
