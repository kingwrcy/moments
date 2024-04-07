import { useEventBus } from '@vueuse/core'
import type { Memo } from './types'

const memoUpdateEvent = useEventBus<Memo>('memoUpdate')
const settingsUpdateEvent = useEventBus<Memo>('settingsUpdate')

export  {memoUpdateEvent,settingsUpdateEvent}