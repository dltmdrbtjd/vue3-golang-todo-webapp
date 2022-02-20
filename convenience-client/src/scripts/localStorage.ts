export function getLocalStorage(key: string) {
    const value = localStorage.getItem(key)
    return value
}

export function setLocalStorage(key: string, value: string) {
    localStorage.setItem(key, value)
}