import type { Arrayable, ResourceCubesTypes } from '~/types'
import type { SelectItem } from '@nuxt/ui'


type CustomSelectItem = SelectItem & {
  label: ResourceCubesTypes,
  value: ResourceCubesTypes
}

export const resourcesConstant: Arrayable<CustomSelectItem>  = [
  {
    label: 'All',
    value: 'All'
  },
  {
    label: 'Material',
    value: 'Material'
  },
  {
    label: 'Energy',
    value: 'Energy'
  },
  {
    label: 'Science',
    value: 'Science'
  },
  {
    label: 'Gold',
    value: 'Gold'
  },
  {
    label: 'Exploration',
    value: 'Exploration'
  },
  {
    label: 'Krystallium',
    value: 'Krystallium'
  }
]
