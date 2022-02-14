export default (envName: string): any => {
    if(!envName) {
        return ""
    }

    return import.meta.env.VITE_APP_BASE_URL;
}