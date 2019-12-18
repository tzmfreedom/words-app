import axios from 'axios'

const getConfig = () => {
  if (process.env.NODE_ENV !== 'production') {
    return {
      auth: {
        username: 'user',
        password: 'pass',
      }
    }
  }
  return {}
}
const baseUrl = process.env.baseUrl;
export default {
  get(url) {
    return axios.get(url, getConfig())
  },
  put(url, data) {
    return axios.put(url, data, getConfig())
  },
  post(url, data) {
    return axios.post(url, data, getConfig())
  },
  delete(url, data) {
    return axios.delete(url, getConfig())
  },
  async findAll() {
    return this.get(`${baseUrl}/api/sentences`)
  },
  async find(id) {
    return this.get(`${baseUrl}/api/sentences/${id}`)
  },
  async create(value) {
    return this.post(`${baseUrl}/api/sentences`, {value: value})
  },
  async update(id, value) {
    return this.put(`${baseUrl}/api/sentences/${id}`, {value: value})
  },
  async destroy(id) {
    return this.delete(`${baseUrl}/api/sentences/${id}`)
  }
}
