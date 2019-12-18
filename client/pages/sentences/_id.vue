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
  import sentence from '~/lib/sentence.js'

  export default {
    validate({ params }) {
      return /^\d+$/.test(params.id)
    },
    async asyncData({ params }) {
      const res = await sentence.find(params.id)
      return {
        sentence: res.data,
      }
    }
  }
</script>

<style>
</style>
