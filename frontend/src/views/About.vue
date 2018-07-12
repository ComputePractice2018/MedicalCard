<template>
  <div class="about">
  <div class="jumbotron">
  <h1 class="display-4">{{'О приложении'}}</h1>
  <hr class="my-4">
  <p>Данное приложение разработано в рамках учебной практики студентами 3 курса направления ИВТ.</p>
  <p>Все совпадения с реальными людьми случайны, команда проекта не несет никакой ответственности за введенные пользователем данные.</p>
  <b-btn v-b-toggle="'collapse1'"  class="btn btn-outline-dark">Команда проекта</b-btn>
  <b-button  class="btn btn-outline-dark" href="https://github.com/orgs/ComputePractice2018/teams/medicalcard">Проект на GitHUB</b-button>
  <p></p>
  <b-collapse  id="collapse1" >
  <b-list-group>
  <h5><b-list-group-item>Золотарева Анастасия <b-badge>Frontend</b-badge></b-list-group-item></h5>
 <h5><b-list-group-item>Минков Олег <b-badge>Backend</b-badge></b-list-group-item></h5>
  <h5><b-list-group-item>Левченко Юлия <b-badge>DevOps</b-badge></b-list-group-item></h5>
</b-list-group>
  </b-collapse>
  </div>
  </div>
</template>
