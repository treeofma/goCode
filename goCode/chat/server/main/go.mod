module main

go 1.23.3

require "message" v0.0.0
replace "message" => "../../common/message" 

require "utils" v0.0.0
replace "utils" => "../utils"

require "process" v0.0.0
replace "process" => "../process"
