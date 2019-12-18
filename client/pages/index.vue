<template>
  <div>
  <h1>Hello World!!!</h1>
  <table>
    <thead>
    <tr>
      <th></th>
      <th>Id</th>
      <th>Value</th>
      <th>CreatedAt</th>
      <th>UpdatedAt</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="sentence in sentences">
      <td>
        <nuxt-link v-bind:to="{name: 'sentences-id', params: { id: sentence.id }}">Detail</nuxt-link> |
        <nuxt-link v-bind:to="{name: 'sentences-sentenceId-edit', params: { sentenceId: sentence.id }}">Edit</nuxt-link> |
        <a href="#" v-on:click.prevent="destroy(sentence)">Delete</a>
      </td>
      <td>{{ sentence.id }}</td>
      <td>{{ sentence.value }}</td>
      <td>{{ $moment(sentence.created_at).format('YYYY/MM/DD hh:mm:ss') }}</td>
      <td>{{ $moment(sentence.updated_at).format('YYYY/MM/DD hh:mm:ss') }}</td>
    </tr>
    </tbody>
  </table>
    <nuxt-link to="/sentences/new">New</nuxt-link>
  </div>
</template>

<script>
  import sentence from '~/lib/sentence.js'

  export default {
    asyncData() {
      return sentence.findAll()
        .then(res => {
          return {
            sentences: res.data.records,
          }
        })
    },
    methods: {
      async destroy(sentence) {
        await sentence.delete(sentence.id);
        this.sentences.splice(this.sentences.indexOf(sentence), 1);
        this.$router.push('/')
      }
    }
  }
</script>

<style>
</style>
