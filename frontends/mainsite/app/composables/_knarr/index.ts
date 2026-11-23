const [usePlayerHandsComposable, _usePlayerHandsStore] = createInjectionState(() => {
  const [showPlayerHands, togglePlayerHands] = useToggle(false)

  return {
    showPlayerHands,
    togglePlayerHands
  }
})

export { usePlayerHandsComposable }

export function usePlayerHandsStore() {
  const store = _usePlayerHandsStore()

  if (!store) {
    throw new Error('usePlayerHandsStore must be used within a provider')
  }

  return store
}

const [useExplorationComposable, _useExplorationStore] = createInjectionState(() => {
  const [showExploration, toggleExploration] = useToggle(true)

  return {
    showExploration,
    toggleExploration
  }
})

export { useExplorationComposable }

export function useExplorationStore() {
  const store = _useExplorationStore()
  if (!store) {
    throw new Error('useExplorationStore must be used within a provider')
  }

  return store
}
