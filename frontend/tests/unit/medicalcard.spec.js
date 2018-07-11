import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import medicalcard from '@/components/medicalcard.vue'

describe('medicalcard.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(medicalcard, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders title', () => {
    const wrapper = shallowMount(medicalcard, {})
    expect(wrapper.text()).to.include('Дата приема')
    expect(wrapper.text()).to.include('ФИО Врача')
    expect(wrapper.text()).to.include('ФИО Медсестры')
    expect(wrapper.text()).to.include('Жалобы')
    expect(wrapper.text()).to.include('Результаты осмотра')
    expect(wrapper.text()).to.include('Диагноз')
    expect(wrapper.text()).to.include('Назначенное лечение')
  })
})
