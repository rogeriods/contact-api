import api from '../api/client'

export interface Contact {
  id?: number
  name: string
  phone: string
}

export const getContacts = async (): Promise<Contact[]> => {
  const res = await api.get('/contacts')
  return res.data
}

export const createContact = async (contact: Contact) => {
  const res = await api.post('/contacts', contact)
  return res.data
}

export const updateContact = async (id: number, contact: Contact) => {
  const res = await api.put(`/contacts/${id}`, contact)
  return res.data
}

export const deleteContact = async (id: number) => {
  await api.delete(`/contacts/${id}`)
}
