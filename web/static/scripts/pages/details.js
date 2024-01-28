import { doFile } from '../components/upload.js'
import loading from '../components/loader.js'

document.querySelector('.upload-area').setAttribute('open', '')

doFile(async (file, metadata) => {
    loading(true)

    const res = await fetch(location.href, {
        method: 'PUT',
        body: file,
        headers: {
            'Content-Type': file.type,
            'File-Private-Url-Expires': metadata.privateUrlExpires,
            'File-Is-Public': metadata.isPublic,
            'File-Auto-Delete-At': metadata.autoDeleteAt
        }
    })

    loading()

    if (res.ok) {
        alert('Updated File successfully')
        location.reload()
    } else {
        const { apiError } = await res.json()
        confirm(apiError.description + ' Do you want to refresh the page?') && location.reload()
    }
})

const deleteBtn = document.querySelector('.delete-btn')
deleteBtn.addEventListener('click', async () => {
    if (confirm('Are you sure?, this action cannot be undone')) {
        loading(true)

        const res = await fetch(location.href, {
            method: 'DELETE'
        })

        loading()
        if (res.ok) {
            const username = location.pathname.split('/')[2]
            location.href = '/dashboard/' + username
        } else {
            const { apiError } = await res.json()
            confirm(apiError.description + ', Do you want to refresh the page?') && location.reload()
        }
    }
})
