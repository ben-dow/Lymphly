import { Title } from "@mantine/core";

export default function Header(){
    return(
        <div className="bg-sky-900 h-full flex  justify-around">
            <div className="flex flex-col justify-center">
                <h1 className="font-sans text-5xl font-light p-5 text-white">Lymphly</h1>
            </div>

            <div className="flex flex-col justify-center h-full">
                <div className="w-25 flex items-center h-full bg-sky-600">
                    <h4 className="text-sky-100 text-xl text-center w-full font-sans font-medium">Search</h4>
                </div>
            </div>
        </div>
    )
}

