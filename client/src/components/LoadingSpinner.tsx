import { Loading } from '@nextui-org/react'
type height = 'full' | 'auto'
export default function LoadingSpinner({ height }: { height?: height }) {
    return (
        <div
            className={`h-${
                height || 'screen'
            } flex items-center justify-center`}
        >
            <Loading size={'xl'} />
        </div>
    )
}
