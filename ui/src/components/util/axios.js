import axios from 'axios'

axios.defaults.baseURL = 'https://api.caioeverest.dev/api'
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.withCredentials = false
axios.defaults.timeout = 10000

export const setDefaultUrl = baseUrl => {
    axios.defaults.baseURL = baseUrl
}

export default axios.create({
    headers: {
        'Cache-Control' : 'no-cache'
    }
})
