import {Box, Button, Divider, Tabs, Text} from '@mantine/core'
import 'radar-sdk-js/dist/radar.css'
import { useEffect, useState } from 'react';
import { PracticeListI } from '../../model/practice';
import { DataDisplay, Map } from '../../components/search/display';
import { Routes, Route, Router, useNavigate } from 'react-router';


export default function Search(){
    return(
        <Box className={"bg-sky-200 p-5 shadow-sm w-full h-full flex flex-col gap-5"}>
            <Box className='flex justify-center'>
                <DataDisplay/>
            </Box>
        </Box>
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



