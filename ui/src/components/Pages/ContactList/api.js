import axios from '../../util/axios'

const ENDPOINT = '/contact'

export const create = (phone, name) => axios.post(ENDPOINT, { phone, name})
    .then(res => res.message)

export const getAll = () => axios.get(ENDPOINT)
    .then(res => res.message)

export const getByPhone = phone => axios.get(`${ENDPOINT}/${phone}`)
    .then(res => res.message)

export const update = (phone, newName) => axios.patch(`${ENDPOINT}/${phone}?new-name=${newName}`)
    .then(res => res.message)

export const clear = phone => axios.delete(`${ENDPOINT}/${phone}`)
    .then(res => res.message)
