import axios from '../../util/axios'

const ENDPOINT = '/contact'

export const create = contactPayload => axios.post(ENDPOINT, contactPayload)
    .then(res => res.data.message)

export const getAll = () => axios.get(ENDPOINT)
    .then(res => res.data)

export const getByPhone = phone => axios.get(`${ENDPOINT}/${phone}`)
    .then(res => res.data)

export const update = data => axios.patch(`${ENDPOINT}/${data.phone}?new_name=${data.name}`)
    .then(res => res.data.message)

export const clear = phone => axios.delete(`${ENDPOINT}/${phone}`)
    .then(res => res.data.message)
