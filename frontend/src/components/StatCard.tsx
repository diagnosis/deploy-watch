export function StatCard({ label, value }: { label: string; value: number }) {
    return (
        <div className="border rounded-md p-4">
            <div className="text-sm text-gray-500">{label}</div>
            <div className="text-2xl font-semibold text-gray-900">{value}</div>
        </div>
    );
}