use std::{io};

#[derive(Debug)]
struct Task {
    name: String,
    complete : bool
}

impl Task {
    fn mark_complete(&mut self) {
        self.complete = !self.complete;
    }
}

pub fn todos_exercise() {
    let mut tasks: Vec<Task> = Vec::new();

    loop {
        print!("Enter a task:");
        
        let mut result = String::with_capacity(20);
        let _input_result = io::stdin()
            .read_line(&mut result)
            .expect("Filed to read line");

        let command = result.trim().to_lowercase();

        if command.is_empty() {
            continue
        }

        if command == "exit" {
            break;
        }

        if command == "list" {
            for (i, task) in tasks.iter().enumerate() {
                print!("[{}]{}. {}", task.complete, i, task.name)
            }
            continue
        }

        if command == "complete" {
            let mut task_id = String::new();
            let _id_result = io::stdin()
                .read_line(&mut task_id)
                .expect("Failed to read line");
            
            match task_id.trim().parse::<usize>() {
                Ok(num) => {
                    tasks[num].mark_complete();
                },
                _ => println!("Invalid task number."),
            }
        }

        let task_name = result.trim().to_string();
        tasks.push(Task {
            name: task_name,
            complete: false
        });

        println!("\nTasks saved: {:?}\n", tasks);
    }
}
