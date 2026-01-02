import { useQuery } from '@tanstack/react-query'
import { appClient } from '../lib/api'

export function useDeploys() {
    const {
        data,
        isLoading,
        isError,
    } = useQuery({
        queryKey: ['deploys'],
        queryFn: () => appClient.getDeployEvents(),
    })

    return { deploys: data ?? [], isLoading, isError }
}