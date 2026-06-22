export const useHelpComposable = createGlobalState(() => {
  const showHelpModal = ref<boolean>(false)
  const toggleShowHelpModal = useToggle(showHelpModal)

  return {
    showHelpModal,
    toggleShowHelpModal
  }
})
