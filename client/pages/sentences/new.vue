<template>
  <div>
    <h1>New Record</h1>
    <form>
      <table>
        <tbody>
        <tr>
          <th>Value</th>
          <td><input type="text" v-model="value" /></td>
        </tr>
        </tbody>
        <button v-on:click.prevent="create" value="">作成</button>
        <nuxt-link to="/">一覧</nuxt-link>
      </table>
    </form>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        value: '',
      }
    },
    methods: {
      async create() {
        const res = await this.$axios.post(`http://localhost:8080/sentences`, { value: this.value }, { auth: { username: 'user', password: 'pass' }})
        this.$router.push({ name: 'sentences-id', params: { id: res.data.id }})
      }
    }
  }
</script>

<style>
</style>
