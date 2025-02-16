import {Box, Divider, Tabs, Text} from '@mantine/core'
import 'radar-sdk-js/dist/radar.css'
import { useEffect, useState } from 'react';
import { PracticeListI } from '../../model/practice';
import { DataDisplay, DataDisplayProps, Map } from '../../components/search/display';
import { Routes, Route, Router, useNavigate } from 'react-router';


export default function Search(){
    const [dataDisplayProps, setDataDisplayProps] = useState<DataDisplayProps>({})

    useEffect(() => {
        fetch("/api/v1/providersearch/practices/all").then((r) =>r.json()).then(j=>{
            const pl: PracticeListI = j
            setDataDisplayProps({
                practiceList: pl,
                mapConfiguration:{
                    RadiusFeature: false
                }
            })
        })
    }, [])
    let navigate = useNavigate()

    return(
        <Box className={"bg-sky-200 p-5 shadow-sm w-full h-full flex flex-col gap-5"}>
            <div className='w-20 items-center flex justify-center h-10 hover:cursor-pointer bg-sky-900 rounded text-white font-sans font-semibold hover:shadow-md hover:bg-sky-800' hidden={window.location.pathname == "/search"} onClick={()=>{navigate("/search")}}>
                Back
            </div>
            <Routes>
                <Route index element={<SearchHome/>}/>
                <Route path="current" element={<SearchByLocation updateDataDisplayProps={setDataDisplayProps}/>}/>
                <Route path="address" element={<SearchByAddress/>}/>
                <Route path="state" element={<SearchByState/>}/>
            </Routes>       

            <Box className='flex justify-center'>
                <DataDisplay {...dataDisplayProps}/>
            </Box>
     
        </Box>
    )
}

interface PracticeUpdaterI{
    updateDataDisplayProps: (props: DataDisplayProps) => void
}

function SearchHome(){
    let navigate = useNavigate()

    return(
        <Box className="flex flex-col gap-5">
            <Box className="text-center">
                <h2 className='text-3xl font-normal font-sans'>Find a Provider</h2>
            </Box>
            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Box onClick={()=>{navigate("current")}} className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By Current Location</Text>
                </Box>

                <Box onClick={()=>{navigate("address")}} className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By Address</Text>
                </Box>

                <Box  onClick={()=>{navigate("state")}} className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>By State</Text>
                </Box>
            </Box>

            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Divider className={"w-full max-w-2xl"} color={"black"}/>
            </Box>
            <Box className='flex flex-row flex-wrap justify-center gap-5'>
                <Box className='h-20 w-50 bg-emerald-900 flex flex-col justify-center text-center shadow-sm hover:shadow-xl hover:cursor-pointer'>
                    <Text c="white" fw={700} size={"lg"}>Browse</Text>
                </Box>
            </Box>


        </Box>            
    )
}

function SearchByLocation(props:PracticeUpdaterI){

    useEffect(()=>{
        navigator.geolocation.getCurrentPosition((pos)=>{
            let lat = pos.coords.latitude
            let long = pos.coords.longitude

            fetch(`/api/v1/providersearch/practices/locate/proximity?lat=${lat}&long=${long}&radius=20`).
                then(res => res.json()).
                then((res)=>{props.updateDataDisplayProps(
                    {
                        practiceList: res,
                        mapConfiguration: {
                            RadiusFeature: true,
                            RadiusOrigin: [long, lat],
                            Radius: 20
                        }
                    }
                )})
            
        })


    }, [])


    return (
        <div>
           <Box className='flex justify-center'>
            </Box>
        </div>
    )
}


function SearchByAddress(){
    return (
        <div>
            Location
        </div>
    )
}


function SearchByState(){
    return (
        <div>
            Location
        </div>
    )
}



