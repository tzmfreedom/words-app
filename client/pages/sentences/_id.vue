<template>
  <div>
    <nuxt-child/>
    <table>
      <tbody>
      <tr>
        <th>Id</th>
        <td>{{ sentence.id }}</td>
      </tr>
      <tr>
        <th>Value</th>
        <td>{{ sentence.value }}</td>
      </tr>
      <tr>
        <th>CreatedAt</th>
        <td>{{ $moment(sentence.created_at).format('YYYY/MM/DD hh:mm:ss') }}</td>
      </tr>
      <tr>
        <th>UpdatedAt</th>
        <td>{{ $moment(sentence.updated_at).format('YYYY/MM/DD hh:mm:ss') }}</td>
      </tr>
      </tbody>
      <nuxt-link to="/">一覧</nuxt-link>
    </table>
  </div>
</template>

<script>
  export default {
    validate({ params }) {
      return /^\d+$/.test(params.id)
    },
    asyncData({ params, $axios }) {
      return $axios.get(`http://@localhost:8080/sentences/${params.id}`, { auth: { username: 'user', password: 'pass' } })
        .then(res => {
          return {
            sentence: res.data,
          }
        })
    }
  }
</script>

<style>
</style>
