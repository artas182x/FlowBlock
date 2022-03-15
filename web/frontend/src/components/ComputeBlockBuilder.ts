import {ComputeNodeBuilder} from "@/components/ComputeNodeBuilder";


export function ComputeBlockBuilder({ MethodData, ChaincodeName }) {
    const Block = new ComputeNodeBuilder("TokenNode", ChaincodeName, MethodData.Name);
    Block.setName(MethodData.Name.split(":")[1])

    Block.addOption("Description", "InputOption")

    MethodData.Arguments.forEach(arg => {
        if (arg.Type === "ts") {
            arg.Value = new Date();
            Block.addOption(arg.Name, "DateOption")
        } else if (arg.Type === "tokenInputs") {
            Block.addOption("Add Input", "AddOption");
            Block.addOption("Remove Input", "AddOption");
        } else {
            arg.Value = ""
            Block.addOption(arg.Name, "InputOption")
        }

    })



    Block.addOutputInterface("Output")
    return Block.build();
}