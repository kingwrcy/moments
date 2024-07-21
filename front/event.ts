import { useEventBus } from '@vueuse/core'

export const memoReloadEvent = useEventBus<string>('memo-list-reload')
export const memoChangedEvent = useEventBus<number>('memo-changed')
