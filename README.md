# LearnGo

## Visual Roadmap :-

```mermaid
flowchart TD
    %% Global Configuration
    %% Class Definitions for Catppuccin Mocha Palette
    classDef phase1 stroke:#cdd6f4,stroke-width:2px,fill:#11111b,color:#cdd6f4;
    classDef phase2 stroke:#89b4fa,stroke-width:2px,fill:#11111b,color:#89b4fa;
    classDef phase3 stroke:#a6e3a1,stroke-width:2px,fill:#11111b,color:#a6e3a1;
    classDef phase4 stroke:#f38ba8,stroke-width:2px,fill:#11111b,color:#f38ba8;
    classDef phase5 stroke:#cba6f7,stroke-width:2px,fill:#11111b,color:#cba6f7;
    
    classDef nodeStyle fill:#1e1e2e,stroke:#45475a,color:#cdd6f4,stroke-width:1px;

    %% Vertical Phase Progression
    M1 --> M2 --> M3 --> M4 --> M5

    %% Subgraphs with horizontal layout internally
    subgraph M1 [Phase 1: Foundations]
        direction LR
        1[1_hello_world] --> 2[2_simple_value] --> 3[3_variables] --> 4[4_constants] --> 5[5_for] --> 6[6_if_else] --> 7[7_switch]
    end

    subgraph M2 [Phase 2: Data & Memory]
        direction LR
        8[8_arrays] --> 9[9_slices] --> 10[10_map] --> 11[11_range] --> 15[15_pointers]
    end

    subgraph M3 [Phase 3: Functional Blocks]
        direction LR
        12[12_function] --> 13[13_variadic] --> 14[14_closures]
    end

    subgraph M4 [Phase 4: Types & Composition]
        direction LR
        16[16_structs] --> 18[18_enums] --> 17[17_interface] --> 19[19_generics]
    end

    subgraph M5 [Phase 5: Concurrency & Systems]
        direction LR
        20[20_goroutines] --> 21[21_channels] --> 22[22_mutex] --> 23[23_files] --> 24[24_package]
    end

    %% Apply Themes
    class M1 phase1;
    class M2 phase2;
    class M3 phase3;
    class M4 phase4;
    class M5 phase5;
    
    class 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24 nodeStyle;