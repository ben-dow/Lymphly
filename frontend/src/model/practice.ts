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

export interface LimitedPracticePracticeListI {
    practices: LimitedPracticeI[]
}

export interface LimitedPracticeI {
    practiceId: string
    name: string
    lattitude: number
    longitude: number
}

export interface ProviderI {
    providerId: string
    practiceId: string
    name: string
    tags: string[]
}

export interface PracticeListI {
    practices: PracticeI[]
}

export interface ProviderListI {
    providers: ProviderI[]
}

