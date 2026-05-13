import { useState } from 'react'
import type { Contact } from '../services/contacts'

interface Props {
  onSubmit: (contact: Contact) => void
  initial?: Contact
}

export default function ContactForm({ onSubmit, initial }: Props) {
  const [name, setName] = useState(initial?.name || '')
  const [phone, setPhone] = useState(initial?.phone || '')

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault()
        onSubmit({ name, phone })
      }}
    >
      <input placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />

      <input placeholder="Phone" value={phone} onChange={(e) => setPhone(e.target.value)} />

      <button type="submit">Save</button>
    </form>
  )
}
