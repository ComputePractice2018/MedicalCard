<template>
  <div class="home">
    <medicalcard title="Медицинская карта пациента ФГБУЗ КБ №8"/>

  </div>
</template>

<script>
// @ is an alias to /src
import medicalcard from '@/components/medicalcard.vue'

export default {
  name: 'home',
  components: {
    medicalcard
  }
}
</script>
