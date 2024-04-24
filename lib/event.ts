import { useEventBus } from "@vueuse/core";
import type { Memo } from "./types";

const memoUpdateEvent = useEventBus<Memo & { index?: number }>("memoUpdate");
const settingsUpdateEvent = useEventBus<Memo>("settingsUpdate");

export { memoUpdateEvent, settingsUpdateEvent };
