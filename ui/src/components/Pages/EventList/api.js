import axios from '../../util/axios'

const ENDPOINT = '/event'

export const create = eventPayload => axios.post(ENDPOINT, eventPayload)
    .then(res => res.data.message)

export const getAll = () => axios.get(ENDPOINT)
    .then(res => res.data)

export const getById = id => axios.get(`${ENDPOINT}/${id}`)
    .then(res => res.data)

export const update = data => axios.patch(`${ENDPOINT}/${data.id}?new_name=${data.name}`)
    .then(res => res.data.message)

export const clear = id => axios.delete(`${ENDPOINT}/${id}`)
    .then(res => res.data.message)
