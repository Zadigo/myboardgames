import type { TableDetails } from "~/types"

export function useFlipSevenLiveGameComposable(ws: ReturnType<typeof useWebSocket<{a: string }>>) {
  const tableDetails = ref<TableDetails>()

  return {
    ws,
    players: ref([]),
    tableDetails
  }
}
