import { useEffect, useState } from 'react'
import { getContacts, createContact, deleteContact } from '../services/contacts'
import type { Contact } from '../services/contacts'
import ContactForm from '../components/ContactForm'

export default function Contacts() {
  const [contacts, setContacts] = useState<Contact[]>([])

  const load = async () => {
    const data = await getContacts()
    setContacts(data)
  }

  useEffect(() => {
    load()
  }, [])

  return (
    <div>
      <h2>Contacts</h2>

      <ContactForm
        onSubmit={async (c) => {
          await createContact(c)
          load()
        }}
      />

      <ul>
        {contacts.map((c) => (
          <li key={c.id}>
            {c.name} - {c.phone}
            <button onClick={() => deleteContact(c.id!)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  )
}
