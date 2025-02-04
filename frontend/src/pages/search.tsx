import {Box, Text} from '@mantine/core'

export default function Search(){
    return(
        <div className="w-full h-full flex justify-center">
            <SearchContainer/>
        </div>
    )
}


function SearchContainer(){
    return(
        <Box className={"bg-emerald-100 p-5 flex flex-col gap-10 shadow-sm w-7xl"}>
            <Box className="text-center">
                <Text size="xl" fw={700}>Find A Provider By:</Text>
            </Box>

            <Box className='flex flex-row flex-wrap justify-between'>
                <Box className='h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Current Location</Text>
                </Box>

                <Box className='h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>State</Text>
                </Box>

                <Box className='h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>City</Text>
                </Box>

                <Box className='h-36 w-36 bg-emerald-800 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Practice</Text>
                </Box>
            </Box>



        </Box>            
    )
}
