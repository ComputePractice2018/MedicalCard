<template>
  <div class="medicalcard">
  <div class="jumbotron">
  <h1 class="display-4">{{ title}}</h1>
  <hr class="my-4">
  <p>Добро пожаловать в медицинскую карту пациента. Здесь вы можете узнать всю информацию о приемах</p>
  <b-btn v-b-toggle="'collapse3'"  class="btn btn-outline-dark">Показать личные данные пациента</b-btn>
  <b-btn  v-if="edit_index == -1" v-b-toggle="'collapse2'"  class="btn btn-outline-dark">Новый прием</b-btn>

  <b-collapse  id="collapse3" >
    <b-card>
      <table  class="table" >
   <thead>
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
    </thead>
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
      </table>
    </b-card>
  </b-collapse>
  <b-collapse  id="collapse2">
    <b-card>
<hr class="my-4">
<h3 v-if="edit_index == -1">Новый прием</h3>
<hr class="my-4">
 <b-card>
<b-form>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-small">Дата приема:</label></b-col>
    <b-col sm="6">
      <b-form-input type="date" v-model="new_appointment.date" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-small">ФИО Врача:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.docname" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-large">ФИО медсестры:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.musename" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-large">Жалобы:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.complaint" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-large">Результаты осмотра:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.checkup" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-large">Диагноз:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.diagnosis" required>
       </b-form-input>
       </b-col>
  </b-row>
  <b-row class="my-1">
    <b-col sm="2"><label for="input-large">Назначенное лечение:</label></b-col>
    <b-col sm="6">
       <b-form-input v-model="new_appointment.treatment" required>
       </b-form-input>
       </b-col>
  </b-row>
  <p><button type="button" class="btn btn-outline-dark" v-if="edit_index == -1" v-on:click="add_new_appointment">Добавить запись</button></p>
  <p><button type="button" class="btn btn-outline-dark" v-if="edit_index > -1" v-on:click="end_of_edition">Сохранить редактирование</button></p>
</b-form>
</b-card>
</b-card>
  </b-collapse>
</div>
    <h3 v-if="error">Ошибка: {{error}}</h3>
    <tr>
    </tr>
  <table  class="table">
   <thead>
   <tr>
    <th> Дата приема</th>
     <th>ФИО Врача</th>
     <th>ФИО Медсестры</th>
    <th>Жалобы</th>
    <th>Результаты осмотра</th>
     <th>Диагноз</th>
     <th>Назначенное лечение</th>
    </tr>
   </thead>
   <tbody>
    <tr v-for="information in appointment" v-bind:key="information.diagnosis">
      <td>{{ information.date }}</td>
      <td>{{ information.docname }}</td>
      <td>{{ information.musename }}</td>
      <td>{{ information.complaint}}</td>
      <td>{{ information.checkup }}</td>
      <td>{{ information.diagnosis}}</td>
      <td>{{ information.treatment }}</td>
      <td><b-button type="button" class="btn btn-outline-warning" v-on:click="edit_appointment(information)">Редактировать запись</b-button></td>
      <td><b-button type="button" class="btn btn-outline-danger" v-on:click="remove_appointment(information)">Удалить запись</b-button></td>
    </tr>
    <tr>
    </tr>
   </tbody>
  </table>
<p></p>
<p></p>
<p></p>
<p></p>
  </div>
</template>

<script>
const axios = require('axios')

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
      error: '',
      appointment: [ ],
      new_appointment:
        {
          'id': 0,
          'date': '',
          'docname': '',
          'musename': '',
          'complaint': '',
          'checkup': '',
          'diagnosis': '',
          'treatment': ''
        }
    }
  },
  mounted: function () {
    this.get_appointments()
  },
  methods: {
    get_appointments: function () {
      const url = '/api/medicalcard/appointments'
      axios.get(url).then(response => {
        this.appointment = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_appointment: function () {
      const url = '/api/medicalcard/appointments'
      axios.post(url, this.new_appointment).then(response => {
        console.log(response)
        this.appointment.push(this.new_appointment)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_appointment: function (item) {
      const url = '/api/medicalcard/appointments/' + this.id
      axios.delete(url).then(response => {
        this.appointment = response.data
      }).catch(response => {
        this.appointment.splice(this.appointment.indexOf(item), 1)
      })
    },
    edit_appointment: function (item) {
      this.edit_index = this.appointment.indexOf(item)
      this.new_appointment = this.appointment[this.edit_index]
    },
    end_of_edition: function () {
      const url = '/api/medicalcard/appointments/' + this.new_appointment.id
      axios.put(url, this.new_appointment).then(response => {
        console.log(response)
        this.edit_index = -1
        this.new_appointment = {
          'id': 0,
          'date': '',
          'docname': '',
          'musename': '',
          'complaint': '',
          'checkup': '',
          'diagnosis': '',
          'treatment': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
