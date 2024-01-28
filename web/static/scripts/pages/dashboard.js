import loading from '../components/loader.js'
import { doFile } from '../components/upload.js'

doFile(async (file, metadata) => {
    loading(true)

    const res = await fetch(location.href, {
        method: 'POST',
        headers: {
            'Content-Type': file.type,
            'File-Auto-Delete-At': metadata.autoDeleteAt,
            'File-Is-Public': metadata.isPublic,
            'File-Name': file.name,
            'File-Private-Url-Expires': metadata.privateUrlExpires,
        },
        body: file,
    })

    loading()
    if (res.ok) {
        alert('Upload successfully')
        location.reload()
    } else {
        const { apiError } = await res.json()
        confirm(apiError.description + ', Do you want to refresh the page?') && location.reload()
    }

})
