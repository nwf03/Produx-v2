import { useAppDispatch } from '../../../../state/hooks'
import { Channel } from '../../../../state/interfaces'
import { setChannel } from '../../../../state/reducers/channelSlice'

export default function ChannelFilters({ channel }: { channel: string }) {
    const channels: Channel = {
        Announcements: { icon: '🎉', color: '#FF9900' },
        Bugs: { icon: '🐞', color: '#DBFF00' },
        Suggestions: { icon: '🙏', color: '#0094FF' },
        Changelogs: { icon: '🔑', color: '#FF4D00' },
    }
    const dispatch = useAppDispatch()
    return (
        <div
            className={
                'flex p-4 gap-8 overflow-x-auto items-center justify-center'
            }
        >
            {Object.keys(channels).map((e) => {
                return (
                    <div
                        key={e}
                        onClick={() => dispatch(setChannel(e))}
                        className={`text-center rounded-2xl p-3 text-sm cursor-pointer ${
                            channel != e && 'text-white'
                        }`}
                        style={{ backgroundColor: channel == e ? 'white' : '' }}
                    >
                        {channels[e].icon} {e}
                    </div>
                )
            })}
        </div>
    )
}
