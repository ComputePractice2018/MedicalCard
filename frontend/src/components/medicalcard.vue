<template>
  <div class="medicalcard">
    <h1>{{ title }}</h1>

  <table cellspacing="15">
   <tr>
     <th>Имя</th>
     <th>Отчество</th>
     <th>Фамилия</th>
     <th>Дата рождения</th>
     <th >Серия/номер паспорта</th>
     <th>Номер страхового свидетельства</th>
     <th>Телефон</th>
     <th>Домашний адрес</th>
    </tr>
    <tr v-for="information in personal_data" v-bind:key="information.passport">
      <td>{{ information.name }}</td>
      <td>{{ information.surname }}</td>
      <td>{{ information.secondname }}</td>
      <td>{{ information.birth}}</td>
      <td>{{ information.passport}}</td>
      <td>{{ information.insurance}}</td>
      <td>{{ information.phone}}</td>
      <td>{{ information.adress}}</td>
    </tr>
    <tr>
    </tr>
  </table>
  <table cellspacing="15">
   <tr>
    <th> Дата приема</th>
     <th>ФИО Врача</th>
     <th>ФИО Медсестры</th>
    <th>Жалобы</th>
    <th>Результаты осмотра</th>
     <th>Диагноз</th>
     <th>Назначенное лечение</th>
    </tr>
    <tr v-for="information in appointment" v-bind:key="information.diagnosis">
      <td>{{ information.date }}</td>
      <td>{{ information.doc_name }}</td>
      <td>{{ information.muse_name }}</td>
      <td>{{ information.complaint}}</td>
      <td>{{ information.checkup }}</td>
      <td>{{ information.diagnosis}}</td>
      <td>{{ information.treatment }}</td>
      <button v-on:click="edit_appointment(information)">Редактировать запись</button>
      <button v-on:click="remove_appointment(information)">Удалить запись</button>

    </tr>
    <tr>
    </tr>
  </table>

<h3 v-if="edit_index == -1">Новый прием</h3>
  <form>
    <p>Дата приема: <input type="text" v-model="new_appointment.date">
    ФИО Врача: <input type="text" v-model="new_appointment.doc_name">
    Фио Медсестры: <input type="text" v-model="new_appointment.muse_name">
    <p></p>
    Жалобы: <input type="text" v-model="new_appointment.complaint">
    Результаты  осмотра: <input type="text" v-model="new_appointment.checkup">
    <p></p>
    Диагноз: <input type="text" v-model="new_appointment.diagnosis">
    Назначенное лечение: <input type="text" v-model="new_appointment.treatment">
  <p><button v-if="edit_index == -1" v-on:click="add_new_appointment">Добавить запись</button></p>
  <p><button v-if="edit_index > -1" v-on:click="make_new_appointment">Сохранить редактирование</button></p>

  </form>
  </div>
</template>

<script>
export default {
  name: 'medicalcard',
  props: {
    title: String
  },
  data: function () {
    return {
      personal_data: [
        {
          'name': 'Анатолий',
          'surname': 'Витальевич',
          'secondname': 'Акапянов',
          'birth': '01.01.1991',
          'passport': '0711661072',
          'insurance': '1234567890123456',
          'adress': 'Обнинск, Ленина 69',
          'phone': '+7-928-8248005'
        }
      ],
      edit_index: -1,
      appointment: [
        {
          'date': '01.07.2018',
          'doc_name': 'Смирнов Василий Данилович',
          'muse_name': 'Прокофьева Вероника Максимовна',
          'complaint': 'Чихает находясь в помещении с цветами',
          'checkup': 'Покраснения слизистой',
          'diagnosis': 'Стоматит',
          'treatment': 'Изолироватть пациента от аллергена'
        },
        {
          'date': '01.07.2018',
          'doc_name': 'Смирнов Василий Данилович',
          'muse_name': 'Прокофьева Вероника Максимовна',
          'complaint': 'Чихает находясь в помещении с цветами',
          'checkup': 'Покраснения слизистой',
          'diagnosis': 'Аллергия',
          'treatment': 'Изолироватть пациента от аллергена'
        }
      ],
      new_appointment:
        {
          'date': '',
          'doc_name': '',
          'muse_name': '',
          'complaint': '',
          'checkup': '',
          'diagnosis': '',
          'treatment': ''
        }
    }
  },
  methods: {
    add_new_appointment: function () {
      this.appointment.push(this.new_appointment)
    },
    remove_appointment: function (item) {
      this.appointment.splice(this.appointment.indexOf(item), 1)
    },
    edit_appointment: function (item) {
      this.edit_index = this.appointment.indexOf(item)
      this.new_appointment = this.appointment[this.edit_index]
    },
    make_new_appointment: function () {
      this.edit_index = -1
      this.new_appointment = {
        'date': '',
        'doc_name': '',
        'muse_name': '',
        'complaint': '',
        'checkup': '',
        'diagnosis': '',
        'treatment': ''
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
