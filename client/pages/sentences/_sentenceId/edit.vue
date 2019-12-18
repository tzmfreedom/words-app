<template>
  <div>
    <h1>Edit Record</h1>
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
  import api from '~/lib/api.js'

  export default {
    async asyncData({ params }) {
      const res = await api.find(params.sentenceId)
      return {
        sentence: res.data,
      }
    },
    methods: {
      async update() {
        const res = await api.update(this.sentence.id, this.sentence.value)
        this.$router.push({ name: 'sentences-id', params: { id: res.data.id }})
      }
    }
  }
</script>

<style>
</style>
