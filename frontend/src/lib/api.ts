import axios from "axios";

const API_URL = import.meta.env.PROD
    ? 'https://backend-cold-bush-2228.fly.dev'  // Production
    : ''  // Development - proxy handles it ✅

const apiClient = axios.create({
    baseURL: API_URL,
    timeout: 10000,
    headers: {'Content-Type': 'application/json'},
    withCredentials: true
})
interface ApiResponse<T> {
    data: T
    correlation_id: string
    timestamp: string
}

export interface User {
    id: string
    github_id: number
    username: string
    email: string
    avatar_url: string
    created_at: string
    updated_at: string
}

export interface DeployEvent {
    id: string
    user_id: string
    repo_name: string
    commit_sha: string
    commit_message: string
    author: string
    branch: string
    status: string
    created_at: string
}

export const appClient = {
    getCurrentUser: async () => {
        const res = await apiClient.get<ApiResponse<{ user: User }>>("/api/me");
        return res.data.data.user;
    },

    getDeployEvents: async () => {
        const res = await apiClient.get<ApiResponse<{ deploys: DeployEvent[] }>>(
            "/api/deploys"
        );
        return res.data.data.deploys;
    },

    getLoginUrl: () => `${API_URL}/auth/github/login`,

    logout: () => {
        window.location.href = `${API_URL}/auth/logout`;
    }
};