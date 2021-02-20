var updateTodo = function(id, isDone){
  $.ajax({
    url: '/graphql?query=mutation+_{update(id:"' + id + '",done:' + isDone + '){id,text,done}}'
  }).done(function(data) {
    console.log(data);
    var updatedTodo = data.data.update;
    if (updatedTodo.done) {
      $('#' + updatedTodo.id).parent().parent().parent().addClass('todo-done');
    } else {
      $('#' + updatedTodo.id).parent().parent().parent().removeClass('todo-done');
    } 
  });
};

var handleTodoList = function(todos) {
  if (!todos.length) {
    $('.todo-list-container').append('<p>There are no tasks for you today</p>');
    return
  } else {
    $('.todo-list-container p').remove();
  }

  $.each(todos, function(i, v) {
    var todoTemplate = $('#todoItemTemplate').html();
    var todo = todoTemplate.replace('{{todo-id}}', v.id);
    todo = todo.replace('{{todo-text}}', v.text);
    todo = todo.replace('{{todo-checked}}', (v.done ? ' checked="checked"' : ''));
    todo = todo.replace('{{todo-done}}', (v.done ? ' todo-done' : ''));

    $('.todo-list-container').append(todo);
    $('#' + v.id).click(function(){
      var id = $(this).prop('id');
      var isDone = $(this).prop('checked');
      updateTodo(id, isDone);
    });
  });
};

var loadTodos = function() {
  $.ajax({
    url: "/graphql?query={list{id,text,done}}"
  }).done(function(data) {
    console.log(data);
    handleTodoList(data.data.list);
  });
};

var addTodo = function(todoText) {
  if (!todoText || todoText === "") {
    alert('Please specify a task');
    return;
  }

  $.ajax({
    url: '/graphql?query=mutation+_{create(text:"' + todoText + '"){id,text,done}}'
  }).done(function(data) {
    console.log(data);
    var todoList = [data.data.create];
    handleTodoList(todoList);
  });
};

$(document).ready(function() {
  $('.todo-add-form').submit(function(e){
    e.preventDefault();
    addTodo($('.todo-add-form #task').val());
    $('.todo-add-form #task').val('');
  });

  loadTodos();
});
