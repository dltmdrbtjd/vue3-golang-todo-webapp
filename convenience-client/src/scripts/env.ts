export default (envName: string): any => {
    if(!envName) {
        return ""
    }
    return import.meta.env[envName];
}