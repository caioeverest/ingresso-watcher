import axios from 'axios'

axios.defaults.baseURL = `${window.location.origin}/api`
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.withCredentials = false
axios.defaults.timeout = 60000

export const setDefaultUrl = baseUrl => {
    axios.defaults.baseURL = baseUrl
}

export default axios.create({
    headers: {
        'Cache-Control' : 'no-cache'
    }
})
