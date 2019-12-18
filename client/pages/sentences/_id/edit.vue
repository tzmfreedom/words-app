<template>
  <div>
    <h1>New Record</h1>
    <form>
      <table>
        <tbody>
        <tr>
          <th>Id</th>
          <td><input type="text" v-model="sentence.id" /></td>
        </tr>
        <tr>
          <th>Value</th>
          <td><input type="text" v-model="sentence.value" /></td>
        </tr>
        <tr>
          <th>CreatedAt</th>
          <td><input type="text" v-model="sentence.created_at" /></td>
        </tr>
        <tr>
          <th>UpdatedAt</th>
          <td><input type="text" v-model="sentence.updated_at" /></td>
        </tr>
        </tbody>
        <button v-on:click.prevent="update" value="">更新</button>
        <nuxt-link to="/">一覧</nuxt-link>
      </table>
    </form>
  </div>
</template>

<script>
  const auth = { auth: { username: 'user', password: 'pass' }};
  export default {
    async asyncData({ params, $axios }) {
      const res = await $axios.get(`http://localhost:8080/sentences/${params.id}`, auth)
      return {
        value: '',
        sentence: res.data,
      }
    },
    methods: {
      async update() {
        const res = await this.$axios.put(`http://localhost:8080/sentences/${this.sentence.id}`, { value: this.value }, auth)
        this.$router.push({ name: 'sentences-id', params: { id: res.data.id }})
      }
    }
  }
</script>

<style>
</style>
