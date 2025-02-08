export interface PracticeI {
    practiceId: string
    name: string
    lattitude: number
    longitude: number
    fullAddress: string
    geoHash: string
    phone: string
    website: string
    state: string
    stateCode: string
    country: string
    countryCode: string
    tags: string[]
}

export interface ProviderI {
    providerId: string
    practiceId: string
    name: string
    tags: string[]
}

export interface PracticeList {
    practices: PracticeI[]
}

export interface ProviderList {
    providers: ProviderI[]
}

